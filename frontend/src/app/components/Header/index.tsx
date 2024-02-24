import Image from "next/image";

import { useTheme } from "../../contexts/Theme";
import ThemeSwitcher from "../ThemeSwitcher";
import UserArea from "../UserArea";

type Props = {
  title: string;
};

const Header = ({ title }: Props) => {
  const { lights, updateTheme } = useTheme();

  const { header, logo } = lights
    ? { header: "border-black", logo: "/dark-logo.png" }
    : { header: "border-white", logo: "/light-logo.png" };

  return (
    <header
      className={`flex items-center justify-between p-6 border-b ${header}`}
    >
      <div className="flex items-center justify-between">
        <Image className="mr-2" width={24} height={24} src={logo} alt="Logo" />
        <h1 className="text-2xl font-bold">{title}</h1>
      </div>
      <div className="flex items-center">
        <nav className="mr-14">
          <UserArea />
        </nav>
        <ThemeSwitcher />
      </div>
    </header>
  );
};

export default Header;
