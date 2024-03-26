import Home from './components/Home';
import { AuthProvider } from './contexts/Auth';
import { ThemeProvider } from './contexts/Theme';

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
