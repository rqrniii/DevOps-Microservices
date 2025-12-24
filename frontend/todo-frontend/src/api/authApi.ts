import axios from "axios";

const authApi = axios.create({
  baseURL: import.meta.env.VITE_AUTH_API || "/api/auth",
});

export default authApi;