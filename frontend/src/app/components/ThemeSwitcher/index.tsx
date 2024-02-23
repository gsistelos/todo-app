import Image from "next/image";
import { useTheme } from "../../contexts/Theme";

const ThemeSwitcher = () => {
  const { lights, updateTheme } = useTheme();

  const { hover, src, alt } = lights
    ? { hover: "hover:bg-gray-200", src: "/dark.png", alt: "Dark icon" }
    : { hover: "hover:bg-gray-800", src: "/light.png", alt: "Light icon" };

  return (
    <button
      className={`p-1 rounded-full ${hover}`}
      onClick={() => updateTheme()}
    >
      <Image width={24} height={24} src={src} alt={alt} />
    </button>
  );
};

export default ThemeSwitcher;
