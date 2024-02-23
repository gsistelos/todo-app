"use client";

import { ThemeContext, ThemeValues } from "./contexts/Theme";
import Header from "./components/Header";

const Home = () => {
  const { lights, updateTheme } = ThemeValues();

  const themeClasses = lights ? "bg-white text-black" : "bg-black text-white";

  return (
    <ThemeContext.Provider value={{ lights, updateTheme }}>
      <main className={`min-h-screen ${themeClasses}`}>
        <Header title="Home" />
        <div className="p-4">
          <h1>Home</h1>
          <span>This is the Home page</span>
        </div>
      </main>
    </ThemeContext.Provider>
  );
};

export default Home;
