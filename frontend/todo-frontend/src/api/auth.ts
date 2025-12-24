import authApi from "./authApi";

export const register = (email: string, password: string) => {
    return authApi.post("/register", { email, password }); 
  };
  
export const login = (email: string, password: string) => {
    return authApi.post("/login", { email, password });  
  };