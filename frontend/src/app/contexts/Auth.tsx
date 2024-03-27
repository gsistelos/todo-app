'use client';

import { createContext, useState, useEffect, useContext } from 'react';

import jwt from 'jsonwebtoken';

const baseURL = 'http://localhost:8080/api';

type User = {
  id: string;
  username: string;
  email: string;
  password: string;
  createdAt: string;
  updatedAt: string;
};

type AuthContextType = {
  user: User | null;
  register: (username: string, email: string, password: string) => Promise<void>;
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
};

export const AuthContext = createContext({} as AuthContextType);

type Props = {
  children: React.ReactNode;
};

export const AuthProvider = ({ children }: Props) => {
  const [user, setUser] = useState<User | null>(null);

  const fetchUser = async (token: string): Promise<void> => {
    const decoded = jwt.decode(token);
    if (!decoded || typeof decoded === 'string') {
      throw new Error('Invalid authentication token');
    }

    const id = decoded.id;

    const response = await fetch(`${baseURL}/users/${id}`, {
      headers: {
        'Content-Type': 'application/json',
        'Authorization': `Bearer ${token}`,
      },
    }).catch(() => {
      throw new Error('Failed to connect to the server');
    });

    if (response.status === 200) {
      const data = await response.json();

      const user = {
        id: data.id,
        username: data.username,
        email: data.email,
        password: data.password,
        createdAt: data.createdAt,
        updatedAt: data.updatedAt,
      };

      setUser(user);
    } else if (response.status === 500) {
      throw new Error('Internal server error');
    } else {
      throw new Error('Invalid authentication token');
    }
  }

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      fetchUser(token).catch(() => {
        localStorage.removeItem('token');
      });
    }
  }, []);

  const register = async (username: string, email: string, password: string): Promise<void> => {
    const response = await fetch(`${baseURL}/users`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        username: username,
        email: email,
        password: password,
      }),
    }).catch(() => {
      throw new Error('Failed to connect to the server');
    });

    if (response.status === 201) {
    } else if (response.status > 399 && response.status < 500) {
      const data = await response.json();
      throw data;
    } else if (response.status == 500) {
      throw new Error('Internal server error');
    } else {
      throw new Error('Unknown error');
    }
  }

  const login = async (email: string, password: string): Promise<void> => {
    const response = await fetch(`${baseURL}/login`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify({
        email: email,
        password: password,
      }),
    }).catch(() => {
      throw new Error('Failed to connect to the server');
    });

    if (response.status === 200) {
      const data = await response.json();

      const token = data.token;

      localStorage.setItem('token', token);

      fetchUser(token).catch(() => {
        localStorage.removeItem('token');
      });
    } else if (response.status > 399 && response.status < 500) {
      const data = await response.json();
      throw data;
    } else if (response.status == 500) {
      throw new Error('Internal server error');
    } else {
      throw new Error('Unknown error');
    }
  }

  const logout = () => {
    localStorage.removeItem('token');
    setUser(null);
  }

  return (
    <AuthContext.Provider value={{ user, register, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

/*
 * 'use client' is needed in order to use useAuth
 */
export const useAuth = () => {
  return useContext(AuthContext);
}
