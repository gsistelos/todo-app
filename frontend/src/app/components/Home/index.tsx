"use client";

import { useTheme } from "../../contexts/Theme";
import Header from "../Header";

const Home = () => {
  const { bgColor, textColor } = useTheme();

  return (
    <main className={`min-h-screen ${textColor} ${bgColor}`}>
      <Header title="Home" />
      <div className="flex flex-col items-center justify-center p-8">
        <h1>Home</h1>
        <span>This is the Home page</span>
      </div>
    </main>
  );
};

export default Home;
