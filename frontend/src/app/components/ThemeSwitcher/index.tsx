import Image from "next/image";

import { useTheme } from "../../contexts/Theme";

const ThemeSwitcher = () => {
  const { lights, updateTheme } = useTheme();

  const { hoverTheme, srcTheme, altTheme } = lights
    ? {
        hoverTheme: "hover:bg-gray-200",
        srcTheme: "/dark.png",
        altTheme: "Dark icon",
      }
    : {
        hoverTheme: "hover:bg-gray-800",
        srcTheme: "/light.png",
        altTheme: "Light icon",
      };

  return (
    <button
      className={`p-1 rounded-full ${hoverTheme}`}
      onClick={() => updateTheme()}
    >
      <Image width={24} height={24} src={srcTheme} alt={altTheme} />
    </button>
  );
};

export default ThemeSwitcher;
