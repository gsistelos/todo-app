import Link from "next/link";
import { useTheme } from "../../contexts/Theme";

const UserArea = () => {
  const { lights } = useTheme();

  const { line, hover } = lights
    ? { hover: "hover:bg-gray-200", line: "border-black" }
    : { hover: "hover:bg-gray-800", line: "border-white" };

  const linkClasses = `px-4 py-2 rounded-full ${hover}`;

  return (
    <ul className="flex">
      <li>
        <Link className={linkClasses} href="/register">
          Register
        </Link>
      </li>
      <div className={`border-l min-h-full mx-3 ${line}`}></div>
      <li>
        <Link className={linkClasses} href="/login">
          Login
        </Link>
      </li>
    </ul>
  );
};

export default UserArea;
