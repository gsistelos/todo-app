'use client';

import { createContext, useContext, useState, useEffect } from 'react';

type ThemeContextType = {
  theme: string;
  toggleTheme: () => void;
  textColor: string;
  bgColor: string;
  bgSecondaryColor: string;
  borderColor: string;
  hoverColor: string;
};

const ThemeContext = createContext<ThemeContextType>({} as ThemeContextType);

type Props = {
  children: React.ReactNode;
};

export const ThemeProvider = ({ children }: Props) => {
  const [theme, setTheme] = useState('dark');

  const toggleTheme = () => {
    if (theme === 'dark') {
      setTheme('light');
      localStorage.setItem('theme', 'light');
    } else {
      setTheme('dark');
      localStorage.setItem('theme', 'dark');
    }
  };

  useEffect(() => {
    const localTheme = localStorage.getItem('theme');
    if (localTheme === 'light') {
      toggleTheme();
    } else if (localTheme !== 'dark') {
      localStorage.setItem('theme', 'dark');
    }
  }, []);

  const value =
    theme === 'dark'
      ? {
        textColor: 'text-white',
        bgColor: 'bg-black',
        bgSecondaryColor: 'bg-gray-800',
        borderColor: 'border-white',
        hoverColor: 'hover:bg-gray-800',
      }
      : {
        textColor: 'text-black',
        bgColor: 'bg-white',
        bgSecondaryColor: 'bg-gray-200',
        borderColor: 'border-black',
        hoverColor: 'hover:bg-gray-200',
      };

  return (
    <ThemeContext.Provider value={{ theme, toggleTheme, ...value }}>
      {children}
    </ThemeContext.Provider>
  );
};

/*
 * 'use client' is needed in order to use useTheme
 */
export const useTheme = () => {
  return useContext(ThemeContext);
};
