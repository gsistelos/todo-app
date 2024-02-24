"use client";

import Header from "./components/Header";
import { ThemeContext, ThemeValues } from "./contexts/Theme";

const Home = () => {
  const { lights, updateTheme } = ThemeValues();

  const mainTheme = lights ? "text-black bg-white" : "text-white bg-black";

  return (
    <ThemeContext.Provider value={{ lights, updateTheme }}>
      <main className={`min-h-screen ${mainTheme}`}>
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
