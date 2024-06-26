'use client';

import { useState } from 'react';

import { useAuth } from '@/app/contexts/Auth';

import FormInput from '../FormInput';
import PasswordSwitcher from '../PasswordSwitcher';

type Props = {
  onClose: () => void;
};

const Login = ({ onClose }: Props) => {
  const [showPassword, setShowPassword] = useState(false);

  const [error, setError] = useState({
    message: '',
    email: '',
    password: '',
  });

  const { login } = useAuth();

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const { email, password } = e.currentTarget;

    login(email.value, password.value)
      .then(() => {
        onClose();
      })
      .catch((error: any) => {
        setError(error);
      });
  };

  return (
    <div className="flex flex-col p-8 gap-4">
      <div className="flex flex-col">
        <h1 className="text-3xl font-semibold">Login</h1>
        <span>Enter your account:</span>
      </div>
      <form className="flex flex-col items-center" onSubmit={handleSubmit}>
        <div className="flex flex-col gap-3">
          {error.message && <span className="block text-red">{error.message}</span>}
          <FormInput error={error.email} placeholder="Email" type="email" name="email" />
          <FormInput error={error.password} placeholder="Password" type={showPassword ? "text" : "password"} name="password" />
          <div className="ml-auto">
            <PasswordSwitcher show={showPassword} onClick={() => setShowPassword(!showPassword)} />
          </div>
        </div>
        <button
          className="bg-secondary px-4 py-2 rounded-full hover:bg-green"
          type="submit"
        >
          Login
        </button>
      </form>
    </div>
  );
};

export default Login;
