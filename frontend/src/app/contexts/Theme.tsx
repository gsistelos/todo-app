import { createContext, useContext, useState, useEffect } from "react";

export const ThemeContext = createContext({
  lights: false,
  updateTheme: () => {},
});

/*
 * "use client" is needed in order to use createTheme
 */
export const createTheme = () => {
  const [lights, setLights] = useState(false);

  useEffect(() => {
    const theme = localStorage.getItem("theme");
    if (theme === "light") {
      setLights(true);
    } else if (theme !== "dark") {
      localStorage.setItem("theme", "dark");
    }
  }, []);

  const updateTheme = () => {
    const newTheme = lights ? "dark" : "light";
    localStorage.setItem("theme", newTheme);
    setLights(!lights);
  };

  return { lights, updateTheme };
};

/*
 * "use client" is needed in order to use useTheme
 */
export function useTheme() {
  const context = useContext(ThemeContext);
  if (!context) {
    throw new Error("useTheme must be used within a ThemeProvider");
  }

  return context;
}
