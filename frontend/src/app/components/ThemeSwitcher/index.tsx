import Image from "next/image";

import { useTheme } from "../../contexts/Theme";

const ThemeSwitcher = () => {
  const { theme, toggleTheme, hoverColor } = useTheme();

  const { src, alt } =
    theme === "dark"
      ? {
          src: "/light.png",
          alt: "Light icon",
        }
      : {
          src: "/dark.png",
          alt: "Dark icon",
        };

  return (
    <button
      className={`p-1 rounded-full ${hoverColor}`}
      onClick={() => toggleTheme()}
    >
      <Image width={24} height={24} src={src} alt={alt} />
    </button>
  );
};

export default ThemeSwitcher;
