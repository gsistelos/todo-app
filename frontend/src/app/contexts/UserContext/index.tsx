import { createContext } from "react";
import { User } from "../../types";
import jwt from "jsonwebtoken";

export const fetchUser = async (): Promise<User | null> => {
  try {
    const token = localStorage.getItem("token");
    if (!token) {
      return null;
    }

    const data = jwt.decode(token);
    if (!data || typeof data === "string") {
      localStorage.removeItem("token");
      return null;
    }

    const response = await fetch("http://localhost:8080/api/users/" + data.id, {
      headers: {
        Authorization: "Bearer " + token,
      },
    });

    if (!response.ok || response.status !== 200) {
      localStorage.removeItem("token");
      return null;
    }

    const user = await response.json();

    user.created_at = new Date(user.created_at);
    user.updated_at = new Date(user.updated_at);

    return user;
  } catch (error) {
    alert("An error occurred while fetching the user: " + error);
    return null;
  }
};

export type UserContextType = {
  user: User | null;
  setUser: (user: User | null) => void;
};

export const UserContext = createContext<UserContextType>({
  user: null,
  setUser: () => {},
});
