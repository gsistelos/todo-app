'use client';

import { createContext, useContext, useState, useCallback, useEffect } from 'react';

type ThemeContextType = {
  theme: string;
  toggleTheme: () => void;
};

const ThemeContext = createContext<ThemeContextType>({} as ThemeContextType);

type Props = {
  children: React.ReactNode;
};

export const ThemeProvider = ({ children }: Props) => {
  const [theme, setTheme] = useState('dark');

  const toggleTheme = useCallback(() => {
    if (theme === 'dark') {
      setTheme('light');
      localStorage.setItem('theme', 'light');
    } else {
      setTheme('dark');
      localStorage.setItem('theme', 'dark');
    }
  }, [theme, setTheme]);

  useEffect(() => {
    const localTheme = localStorage.getItem('theme');
    if (localTheme === 'light') {
      toggleTheme();
    } else if (localTheme !== 'dark') {
      localStorage.setItem('theme', 'dark');
    }
  }, [toggleTheme]);

  return (
    <ThemeContext.Provider value={{ theme, toggleTheme }}>
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
