## Updated the Swagger API doc and added test cases for all the new changes that I made on the BackEnd codebase.

## Task No 1: Enable Searching by Task and Status

### Objective
Enhance the existing `GET /api/v1/todos` endpoint to allow users to filter the list of todos based on task name and status values.

### Changes Made

1. **Endpoint Modification**:
   - Updated the `GET /api/v1/todos` endpoint to accept two optional query parameters: `task` and `status`.

2. **Implementation Details**:
   - The modified endpoint can now handle requests that include the `task` and `status` parameters. This allows users to filter todos effectively.

### Example Request
To fetch todos filtered by task name and status, the following request can be made:

```
GET /api/v1/todos?task={taskName}&status={statusValue}
```

- **Parameters**:
  - `task`: A string that allows for partial matching on the task name.
  - `status`: A string that filters todos by their current status (e.g., `created`, `processing`, `done`).

### Conclusion
The enhancements to the `GET /api/v1/todos` endpoint allow for greater flexibility in retrieving todos based on specific search criteria, improving the overall usability of the TODO application.

## Task 2: Add Task Sorting Feature by Priority
### Objective
Implement a sorting feature for tasks in the TODO application, allowing users to view tasks sorted by priority, ensuring that higher-priority tasks are displayed at the top.

### Changes Made
1. **Model Update**:
    - Added a new field to the Todo model to represent the priority of each task. This field can take values such as low, medium, or high.

2. **Sorting Implementation**:
    - Implemented functionality in the GET /api/v1/todos endpoint to sort tasks based on their priority. Tasks are now sorted in descending order, ensuring that higher-priority tasks appear first in the list.

### Conclusion
The addition of a sorting feature by priority significantly enhances the usability of the TODO application, allowing users to focus on high-priority tasks first and improving overall task management.


## Task No 3: Display Completed Tasks Separately from Incomplete Tasks Objective

### Objective
Enhance the TODO application by displaying completed tasks (done tasks) separately from incomplete tasks. This allows users to easily differentiate between tasks that are still pending and those that have been completed.


### Changes Made

1. **Separation of Tasks**:
    - Modified the UI to ensure that tasks with a status of done are displayed in a separate section at the bottom of the task list.

2. **Rendering Logic**:
   - Implemented logic in the component to filter and render tasks based on their status, ensuring that the UI reflects this separation effectively.


### Conclusion
By separating the display of completed tasks from incomplete tasks, users can more easily track their progress and focus on outstanding work. This change enhances the overall usability of the TODO application.
