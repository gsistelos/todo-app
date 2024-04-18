'use client';

import { createContext, useState, useEffect, useContext } from 'react';

import { toast } from 'react-toastify';

import { fetchRegister, fetchLogin, fetchUser } from '@/api';
import { User } from '@/types';

type AuthContextType = {
  token: string | null;
  user: User | null;
  register: (username: string, email: string, password: string, confirmPassword: string) => Promise<void>;
  login: (email: string, password: string) => Promise<void>;
  logout: () => void;
};

export const AuthContext = createContext({} as AuthContextType);

type Props = {
  children: React.ReactNode;
};

export const AuthProvider = ({ children }: Props) => {
  const [token, setToken] = useState<string | null>(null);
  const [user, setUser] = useState<User | null>(null);

  useEffect(() => {
    const token = localStorage.getItem('token');
    if (token) {
      setToken(token);

      fetchUser({ token })
        .then((user) => {
          setUser(user);
        })
        .catch(() => {
          localStorage.removeItem('token');
          toast.error('Failed to retrieve user data');
        });
    }
  }, []);

  const register = async (username: string, email: string, password: string, confirmPassword: string): Promise<void> => {
    await fetchRegister(
      username,
      email,
      password,
      confirmPassword
    );

    toast.success('Registered successfully!');
  }

  const login = async (email: string, password: string): Promise<void> => {
    const token = await fetchLogin(email, password);
    setToken(token);

    const user = await fetchUser({ token })
      .catch(() => {
        toast.error('Failed to retrieve user data');
      });
    setUser(user);

    localStorage.setItem('token', token);
    toast.success('Logged in successfully!');
  }

  const logout = () => {
    localStorage.removeItem('token');
    setToken(null);
    setUser(null);

    toast.success('Logged out');
  }

  return (
    <AuthContext.Provider value={{ token, user, register, login, logout }}>
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
