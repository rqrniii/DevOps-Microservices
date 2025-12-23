import axios from "axios";

const authApi = axios.create({
  baseURL: import.meta.env.VITE_AUTH_API || "http://localhost:8080",
});

export default authApi;