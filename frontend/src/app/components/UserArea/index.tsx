import { useTheme } from "../../contexts/Theme";
import { useState } from "react";
import Modal from "../Modal";
import Register from "../Register";
import Login from "../Login";

const UserArea = () => {
  const { lights } = useTheme();

  const { line, hover } = lights
    ? { hover: "hover:bg-gray-200", line: "border-black" }
    : { hover: "hover:bg-gray-800", line: "border-white" };

  const linkClasses = `px-4 py-2 rounded-full ${hover}`;

  const [isLoginOpen, setIsLoginOpen] = useState(false);
  const [isRegisterOpen, setIsRegisterOpen] = useState(false);

  return (
    <ul className="flex">
      <li>
        <button className={linkClasses} onClick={() => setIsRegisterOpen(true)}>
          Register
        </button>
        {isRegisterOpen && (
          <Modal onClose={() => setIsRegisterOpen(false)}>
            <Register />
          </Modal>
        )}
      </li>
      <div className={`border-l min-h-full mx-3 ${line}`}></div>
      <li>
        <button className={linkClasses} onClick={() => setIsLoginOpen(true)}>
          Login
        </button>
        {isLoginOpen && (
          <Modal onClose={() => setIsLoginOpen(false)}>
            <Login />
          </Modal>
        )}
      </li>
    </ul>
  );
};

export default UserArea;
