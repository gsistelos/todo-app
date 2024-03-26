'use client';

import { useState } from 'react';

import { useAuth } from '@/app/contexts/Auth';
import { useTheme } from '@/app/contexts/Theme';

import FormInput from '../FormInput';
import PasswordSwitcher from '../PasswordSwitcher';

const Register = () => {
  const [showPassword, setShowPassword] = useState(false);

  const [formError, setFormError] = useState({
    message: '',
    username: '',
    email: '',
    password: '',
    confirmPassword: '',
  });

  const { register } = useAuth();


  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const { username, email, password, confirmPassword } = e.currentTarget;

    if (password.value !== confirmPassword.value) {
      setFormError({
        ...formError,
        confirmPassword: 'Passwords do not match',
      });
      return;
    }

    register(username.value, email.value, password.value)
      .catch((error: any) => {
        setFormError(error);
      });
  };

  const { bgColor, bgSecondaryColor, borderColor } = useTheme();

  return (
    <div className={`flex flex-col ${bgColor} p-8 border ${borderColor} rounded-lg gap-4`} >
      <div className="flex flex-col">
        <h1 className="text-3xl font-semibold">Register</h1>
        <span>Create your account:</span>
      </div>
      <form className="flex flex-col items-center" onSubmit={handleSubmit}>
        <div className="flex flex-col gap-3">
          {formError.message && <span className="block text-red-500">{formError.message}</span>}
          <FormInput error={formError.username} placeholder="Username" type="text" name="username" />
          <FormInput error={formError.email} placeholder="Email" type="email" name="email" />
          <FormInput error={formError.password} placeholder="Password" type={showPassword ? "text" : "password"} name="password" />
          <FormInput error={formError.confirmPassword} placeholder="Confirm Password" type={showPassword ? "text" : "password"} name="confirmPassword" />
          <div className="ml-auto">
            <PasswordSwitcher show={showPassword} onClick={() => setShowPassword(!showPassword)} />
          </div>
        </div>
        <button
          className={`${bgSecondaryColor} px-4 py-2 rounded-full hover:bg-green-500`}
          type="submit"
        >
          Resgister
        </button>
      </form>
    </div>
  );
};

export default Register;
