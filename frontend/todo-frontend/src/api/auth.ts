import authApi from "./authApi";

export const register = (email: string, password: string) => {
    return authApi.post("/auth/register", { email, password });
  };
  
  export const login = (email: string, password: string) => {
    return authApi.post("/auth/login", { email, password });
  };
  