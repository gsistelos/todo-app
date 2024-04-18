import { baseUrl } from '@/api';

export const fetchRegister = async (username: string, email: string, password: string, confirm_password: string): Promise<void> => {
  const response = await fetch(`${baseUrl}/auth/register`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      username,
      email,
      password,
      confirm_password,
    }),
  }).catch(() => {
    throw new Error('Failed to connect to the server');
  });

  if (response.status === 400 || response.status === 409) {
    const data = await response.json();
    throw data;
  } else if (response.status !== 201) {
    throw new Error(response.statusText);
  }
}

export const fetchLogin = async (email: string, password: string): Promise<string> => {
  const response = await fetch(`${baseUrl}/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({
      email,
      password,
    }),
  }).catch(() => {
    throw new Error('Failed to connect to the server');
  });

  if (response.status === 400) {
    const data = await response.json();
    throw data;
  } else if (response.status === 401) {
    throw new Error('Invalid email or password');
  } else if (response.status !== 200) {
    throw new Error(response.statusText);
  }

  const data = await response.json();
  return data.token;
}
