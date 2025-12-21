import todoApi from "./todoApi";

export const getTodos = () => {
    return todoApi.get("/todos", {
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      }
    });
  };
  
  export const addTodo = (task: string) => {
    return todoApi.post(
      "/todos",
      { task },
      {
        headers: {
          Authorization: `Bearer ${localStorage.getItem("token")}`,
        },
      }
    );
  };