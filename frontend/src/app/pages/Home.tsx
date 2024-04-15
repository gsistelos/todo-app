'use client';

import { ToastContainer } from 'react-toastify';

import 'react-toastify/dist/ReactToastify.css';

import { useAuth } from '@/app/contexts/Auth';
import { useTheme } from '@/app/contexts/Theme';

import Header from '../components/Header';

const Home = () => {
  const { user } = useAuth();

  const { theme } = useTheme();

  return (
    <main className={`${theme} bg-primary min-h-screen text-contrast`}>
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
      <div id="modal-root"></div>
      <ToastContainer
        position="bottom-right"
        autoClose={5000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover={false}
        theme={theme}
      />
    </main>
  );
};

export default Home;
