import { AuthProvider } from './contexts/Auth';
import { ThemeProvider } from './contexts/Theme';
import Home from './pages/Home';

const App = () => {
  return (
    <AuthProvider>
      <ThemeProvider>
        <Home />
      </ThemeProvider>
    </AuthProvider>
  );
};

export default App;
