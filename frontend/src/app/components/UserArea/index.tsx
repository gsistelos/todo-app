import { useState } from 'react';

import { useTheme } from '@/app/contexts/Theme';
import { useAuth } from '@/app/contexts/Auth';

import Login from '../Login';
import Modal from '../Modal';
import Register from '../Register';

const UserArea = () => {
  const { user, logout } = useAuth();

  const [isLoginOpen, setIsLoginOpen] = useState(false);
  const [isRegisterOpen, setIsRegisterOpen] = useState(false);

  const { hoverColor } = useTheme();

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

  return user ? (
    <ul className="flex gap-1">
      <li>
        <Button onClick={() => { }}>Profile</Button>
      </li>
      <li>
        <Button onClick={logout}>Logout</Button>
      </li>
    </ul>
  ) : (
    <ul className="flex gap-1">
      <li>
        <Button onClick={() => setIsRegisterOpen(true)}>Register</Button>
        {isRegisterOpen && (
          <Modal onClose={() => setIsRegisterOpen(false)}>
            <Register />
          </Modal>
        )}
      </li>
      <li>
        <Button onClick={() => setIsLoginOpen(true)}>Login</Button>
        {isLoginOpen && (
          <Modal onClose={() => setIsLoginOpen(false)}>
            <Login />
          </Modal>
        )}
      </li>
    </ul>
  )
}

export default UserArea;
