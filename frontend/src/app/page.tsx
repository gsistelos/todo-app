import Home from "./components/Home";
import { ThemeProvider } from "./contexts/Theme";

const App = () => {
  return (
    <ThemeProvider>
      <Home />
    </ThemeProvider>
  );
};

export default App;
