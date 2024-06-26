'use client';

import { useState } from 'react';

import { useAuth } from '@/app/contexts/Auth';

import Login from '../Login';
import Modal from '../Modal';
import Register from '../Register';

const UserArea = () => {
  const { user, logout } = useAuth();

  const [isLoginOpen, setIsLoginOpen] = useState(false);
  const [isRegisterOpen, setIsRegisterOpen] = useState(false);

  type Props = {
    children: React.ReactNode;
    onClick: () => void;
  };

  const Button = ({ children, onClick }: Props) => {
    return (
      <button
        className="px-4 py-2 rounded-full hover:bg-secondary"
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
            <Register onClose={() => setIsRegisterOpen(false)} />
          </Modal>
        )}
      </li>
      <li>
        <Button onClick={() => setIsLoginOpen(true)}>Login</Button>
        {isLoginOpen && (
          <Modal onClose={() => setIsLoginOpen(false)}>
            <Login onClose={() => setIsLoginOpen(false)} />
          </Modal>
        )}
      </li>
    </ul>
  )
}

export default UserArea;
