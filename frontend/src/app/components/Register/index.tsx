import { useState } from "react";

import Input from "../FormField";
import PasswordSwitcher from "../PasswordSwitcher";

type FormData = {
  username: string;
  email: string;
  password: string;
  confirmPassword: string;
};

const Register = () => {
  const [showPassword, setShowPassword] = useState(false);

  const [formData, setFormData] = useState<FormData>({
    username: "",
    email: "",
    password: "",
    confirmPassword: "",
  });

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value });
  };

  const handleSubmit = (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    console.log(formData);
  };

  return (
    <div>
      <div className="flex flex-col mb-4">
        <h1 className="text-3xl font-semibold">Register</h1>
        <span>Create your account:</span>
      </div>
      <form className="flex flex-col items-center" onSubmit={handleSubmit}>
        <div className="mb-4">
          <Input
            className="flex flex-col mb-2"
            label="Username:"
            type="text"
            name="username"
            onChange={handleChange}
          />
          <Input
            className="flex flex-col mb-2"
            label="Email:"
            type="email"
            name="email"
            onChange={handleChange}
          />
          <Input
            className="flex flex-col mb-2"
            label="Password:"
            type={showPassword ? "text" : "password"}
            name="password"
            onChange={handleChange}
          />
          <Input
            className="flex flex-col mb-1"
            label="Confirm password:"
            type={showPassword ? "text" : "password"}
            name="confirmPassword"
            onChange={handleChange}
          />
          <PasswordSwitcher
            show={showPassword}
            onClick={() => setShowPassword(!showPassword)}
          />
        </div>
        <button
          className="px-4 py-2 rounded-full hover:bg-green-500"
          type="submit"
        >
          Register
        </button>
      </form>
    </div>
  );
};

export default Register;
