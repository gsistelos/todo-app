'use client';

import { useState } from 'react';

import { useAuth } from '@/app/contexts/Auth';

import FormInput from '../FormInput';
import PasswordSwitcher from '../PasswordSwitcher';

type Props = {
  onClose: () => void;
};

const Register = ({ onClose }: Props) => {
  const [showPassword, setShowPassword] = useState(false);
  const [error, setError] = useState({
    message: '',
    username: '',
    email: '',
    password: '',
    confirm_password: '',
  });

  const { register } = useAuth();

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const { username, email, password, confirmPassword } = e.currentTarget;

    register(username.value, email.value, password.value, confirmPassword.value)
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
        <h1 className="text-3xl font-semibold">Register</h1>
        <span>Create your account:</span>
      </div>
      <form className="flex flex-col items-center" onSubmit={handleSubmit}>
        <div className="flex flex-col gap-3">
          {error.message && <span className="block text-red">{error.message}</span>}
          <FormInput error={error.username} placeholder="Username" type="text" name="username" />
          <FormInput error={error.email} placeholder="Email" type="email" name="email" />
          <FormInput error={error.password} placeholder="Password" type={showPassword ? "text" : "password"} name="password" />
          <FormInput error={error.confirm_password} placeholder="Confirm password" type={showPassword ? "text" : "password"} name="confirmPassword" />
          <div className="ml-auto">
            <PasswordSwitcher show={showPassword} onClick={() => setShowPassword(!showPassword)} />
          </div>
        </div>
        <button
          className="bg-secondary px-4 py-2 rounded-full hover:bg-green"
          type="submit"
        >
          Resgister
        </button>
      </form>
    </div>
  );
};

export default Register;
