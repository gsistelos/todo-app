import { useState } from "react";

import { useTheme } from "../../contexts/Theme";
import Login from "../Login";
import Modal from "../Modal";
import Register from "../Register";

const UserArea = () => {
  const [isLoginOpen, setIsLoginOpen] = useState(false);
  const [isRegisterOpen, setIsRegisterOpen] = useState(false);

  const { borderColor, hoverColor } = useTheme();

  type Props = {
    children: React.ReactNode;
    onClick: () => void;
  };

  const Button = ({ children, onClick }: Props) => {
    return (
      <button
        className={`px-4 py-2 rounded-full ${hoverColor}`}
        onClick={onClick}
      >
        {children}
      </button>
    );
  };

  return (
    <ul className="flex">
      <li>
        <Button onClick={() => setIsRegisterOpen(true)}>Register</Button>
        {isRegisterOpen && (
          <Modal onClose={() => setIsRegisterOpen(false)}>
            <Register />
          </Modal>
        )}
      </li>
      <div className={`min-h-full mx-3 border-l ${borderColor}`}></div>
      <li>
        <Button onClick={() => setIsLoginOpen(true)}>Login</Button>
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
