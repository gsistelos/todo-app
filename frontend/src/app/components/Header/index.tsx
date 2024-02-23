import Link from "next/link";
import { useTheme } from "../../contexts/Theme";
import { link } from "fs";

type Props = {
  title: string;
};

const Header = ({ title }: Props) => {
  const { lights, updateTheme } = useTheme();

  const { headerTheme, logo, linkTheme, lineTheme, buttonIcon, buttonAlt } =
    lights
      ? {
          headerTheme: "border-black",
          logo: "/dark-logo.png",
          linkTheme: "hover:bg-gray-200",
          lineTheme: "bg-black",
          buttonIcon: "/dark.png",
          buttonAlt: "Dark icon",
        }
      : {
          headerTheme: "border-white",
          logo: "/light-logo.png",
          linkTheme: "hover:bg-gray-800",
          lineTheme: "bg-white",
          buttonIcon: "/light.png",
          buttonAlt: "Light icon",
        };

  const linkClasses = `px-4 py-3 rounded-full ${linkTheme}`;

  return (
    <header
      className={`flex items-center justify-between p-6 border-b ${headerTheme}`}
    >
      <div className="flex items-center justify-between">
        <img className="mr-2" src={logo} alt="Logo" />
        <h1 className="text-2xl font-bold">{title}</h1>
      </div>
      <div className="flex justify-between">
        <nav className="mr-14">
          <ul className="flex">
            <li>
              <Link className={linkClasses} href="/register">
                Register
              </Link>
            </li>
            <li className={`w-px mx-3 ${lineTheme}`}></li>
            <li>
              <Link className={linkClasses} href="/login">
                Login
              </Link>
            </li>
          </ul>
        </nav>
        <button onClick={() => updateTheme()}>
          <img src={buttonIcon} alt={buttonAlt} />
        </button>
      </div>
    </header>
  );
};

export default Header;
