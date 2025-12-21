import axios from "axios";

const todoApi = axios.create({
  baseURL: "http://localhost:8081",
});

export default todoApi;