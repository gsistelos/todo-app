import jwt from 'jsonwebtoken';

import { baseUrl } from './index';

const getIdFromToken = (token: string): string => {
  const decoded = jwt.decode(token);
  if (!decoded || typeof decoded === 'string') {
    throw new Error('Invalid authentication token');
  }

  return decoded.id;
}

export const fetchUser = async ({ method, token, body }: { method?: string, token?: string, body?: object }): Promise<any> => {
  const id = token ? getIdFromToken(token) : undefined;
  const url = `${baseUrl}/users${id ? `/${id}` : ''}`;

  const headers = new Headers();
  headers.append('Content-Type', 'application/json');

  if (token) {
    headers.append('Authorization', `Bearer ${token}`);
  }

  const response = await fetch(url, {
    method,
    headers,
    body: body ? JSON.stringify(body) : undefined,
  }).catch(() => {
    throw new Error('Failed to connect to the server');
  });

  if (response.status === 400) {
    throw await response.json();
  } else if (response.status === 404) {
    throw new Error('User not found');
  } else if (response.status !== 200) {
    throw new Error(response.statusText);
  }

  return await response.json();
}
