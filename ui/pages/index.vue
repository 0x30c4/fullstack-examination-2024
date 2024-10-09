
<template>
  <div class="todo-main">
    <h1>TODO List</h1>
    <div v-if="statusMessage" class="status-message">{{ statusMessage }}</div>
    <div class="input-group">
      <input v-model="newTask" placeholder="Enter new task" @keyup.enter="addTodo">
      <select v-model="newTaskPriority">
        <option value="low">Low</option>
        <option value="medium">Medium</option>
        <option value="high">High</option>
      </select>
      <button @click="addTodo">Add</button>
    </div>

    <!-- Add search fields -->
    <div class="search-group">
      <input v-model="searchTask" placeholder="Search by task name" @input="searchTodos">
      <select v-model="searchStatus" @change="searchTodos">
        <option value="">Select</option>
        <option value="created">Created</option>
        <option value="processing">On Going</option>
        <option value="done">Done</option>
      </select>
    </div>

    <div class="todo-columns">
      <div class="todo-column">
        <h2>On Going Tasks</h2>
        <div v-if="ongoingTodos.length > 0">
          <div v-for="todo in ongoingTodos" :key="todo.ID" class="todo-item">
            <input
              v-if="todo.isEditing"
              v-model="todo.Task"
              class="edit-input"
              @blur="editTodo(todo)"
              @keyup.enter="editTodo(todo)"
            >
            <span v-else :class="{ 'done-task': todo.Status === 'done' }" @click="enableEdit(todo)">
              {{ todo.Task }}
            </span>
            <div class="buttons">
              <button :class="{ 'done': todo.Status === 'done' }" @click="updateStatus(todo)">
                ✔️
              </button>
              <button class="delete-button" @click="deleteTodo(todo.ID)">🗑️</button>
            </div>
          </div>
        </div>
        <div v-else>
          <p>進行中のタスクがありません。</p>
        </div>
      </div>

      <div class="todo-column">
        <h2>Done Tasks</h2>
        <div v-if="doneTodos.length > 0">
          <div v-for="todo in doneTodos" :key="todo.ID" class="todo-item">
            <span class="done-task">{{ todo.Task }}</span>
            <div class="buttons">
              <button :class="{ 'done': todo.Status === 'done' }" @click="updateStatus(todo)">
                ✔️
              </button>
              <button class="delete-button" @click="deleteTodo(todo.ID)">🗑️</button>
            </div>
          </div>
        </div>
        <div v-else>
          <p>完了したタスクがありません。</p>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      newTask: '',
      newTaskPriority: 'low', // Priority for new tasks
      todos: [],
      statusMessage: '',
      searchTask: '',
      searchStatus: '', // Search filters
    };
  },
  computed: {
    ongoingTodos() {
      return this.todos.filter(todo => todo.Status !== 'done');
    },
    doneTodos() {
      return this.todos.filter(todo => todo.Status === 'done');
    }
  },
  mounted() {
    this.searchTodos();
  },
  methods: {
    async fetchTodos() {
      try {
        const response = await fetch(`/api/v1/todos`);
        if (!response.ok) throw new Error(`Failed to get todo list. statusCode: ${response.status}`);
        const data = await response.json();
        this.todos = data.data;
      } catch (error) {
        console.error(error);
        this.statusMessage = 'タスクの取得に失敗しました';
      }
    },
    async addTodo() {
      if (this.newTask.trim() === '') return;

      try {
        const response = await fetch('/api/v1/todos', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            task: this.newTask,
            priority: this.newTaskPriority, // Include priority
            Status: 'created'
          })
        });

        if (!response.ok) throw new Error(`Failed to create todo. statusCode: ${response.status}`);

        const data = await response.json();
        this.todos.push(data.data);
        this.newTask = '';
        this.newTaskPriority = 'low'; // Reset priority
        this.statusMessage = 'タスクが追加されました';
      } catch (error) {
        console.error('Error creating todo:', error);
        this.statusMessage = 'タスクの作成に失敗しました';
      }
    },
    async searchTodos() {
      try {
        // Fetch todos with search filters (task and status)
        const response = await fetch(`/api/v1/todos?task=${this.searchTask}&status=${this.searchStatus}&sortBy=priority`);
        if (!response.ok) throw new Error(`Failed to get todo list. statusCode: ${response.status}`);

        const data = await response.json();
        this.todos = data.data;
      } catch (error) {
        console.error('Error fetching todos:', error);
        this.statusMessage = 'タスクの取得に失敗しました';
      }
    },
    enableEdit(todo) {
      todo.isEditing = true;
    },
    async editTodo(todo) {
      todo.isEditing = false;

      try {
        const response = await fetch(`/api/v1/todos/${todo.ID}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            task: todo.Task
          })
        });

        if (!response.ok) throw new Error(`Failed to edit todo. statusCode: ${response.status}`);
        this.statusMessage = 'タスクが編集されました';
      } catch (error) {
        console.error('Error editing todo:', error);
        this.statusMessage = 'タスクの編集に失敗しました';
      }
    },
    async updateStatus(todo) {
      try {
        const response = await fetch(`/api/v1/todos/${todo.ID}`, {
          method: 'PUT',
          headers: {
            'Content-Type': 'application/json',
          },
          body: JSON.stringify({
            Status: todo.Status === 'done' ? 'created' : 'done'
          })
        });

        if (!response.ok) throw new Error(`Failed to update todo Status. statusCode: ${response.status}`);

        todo.Status = todo.Status === 'done' ? 'created' : 'done';
        this.statusMessage = 'タスクのステータスが変更されました';
      } catch (error) {
        console.error('Error updating todo Status:', error);
        this.statusMessage = 'ステータスの更新に失敗しました';
      }
    },
    async deleteTodo(id) {
      try {
        const response = await fetch(`/api/v1/todos/${id}`, {
          method: 'DELETE',
          headers: {
            'Content-Type': 'application/json',
          },
        });

        if (!response.ok) throw new Error(`Failed to delete todo. statusCode: ${response.status}`);

        this.todos = this.todos.filter(todo => todo.ID !== id);
        this.statusMessage = 'タスクが削除されました';
      } catch (error) {
        console.error('Error deleting todo:', error);
        this.statusMessage = 'タスクの削除に失敗しました';
      }
    }
  }
};
</script>

<style scoped>
/* Add relevant styles */
</style><style scoped>
.todo-main {
  max-width: 400px;
  margin: 20px auto;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
  background-color: #fff;
}

.input-group {
  display: flex;
  margin-bottom: 20px;
}

input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
  margin-right: 10px;
}

button {
  padding: 10px;
  border: none;
  background-color: #333;
  color: #fff;
  border-radius: 4px;
  cursor: pointer;
}

.todo-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}

.buttons button {
  background-color: #f0f0f0;
  color: #333;
  margin-left: 5px;
  border-radius: 4px;
  padding: 5px 10px;
  transition: background-color 0.3s, color 0.3s;
}

.buttons button.done {
  background-color: #333;
  color: #fff;
}

.buttons button.done::before {
  color: #fff;
}

.buttons button.delete-button {
  color: white;
}

.status-message {
  margin-bottom: 20px;
  padding: 10px;
  background-color: #f0f0f0;
  border-radius: 4px;
}

.done-task {
  text-decoration: line-through;
}

.edit-input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ddd;
  border-radius: 4px;
}
.todo-columns {
  display: flex;
  justify-content: space-between;
}

.todo-column {
  flex: 1;
  margin-right: 20px; /* Add spacing between columns */
}

.todo-column:last-child {
  margin-right: 0; /* Remove margin for the last column */
}

.todo-column h2 {
  margin-top: 0;
}
</style>
