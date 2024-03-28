'use client';

import { useAuth } from '@/app/contexts/Auth';
import { useTheme } from '@/app/contexts/Theme';

import Header from '../components/Header';

const Home = () => {
  const { user } = useAuth();

  const { theme } = useTheme();

  const themeClass = theme === 'dark' ? 'theme-dark' : 'theme-light';

  return (
    <main className={`${themeClass} bg-primary min-h-screen text-contrast`}>
      <Header title="Home" />
      <div className="flex flex-col items-center justify-center p-8">
        <h1>Home</h1>
        {user ? (
          <div>
            <h2>Welcome, {user.username}!</h2>
          </div>
        ) : (
          <div>
            <h2>Welcome, guest!</h2>
          </div>
        )}
      </div>
    </main>
  );
};

export default Home;
