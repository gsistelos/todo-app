import styles from "./page.module.css";
import Header from "../components/Header";
import LoginForm from "../components/LoginForm";

const Login = () => {
  return (
    <main className={styles.main}>
      <Header title="Login" subtitle="Enter your account" />
      <LoginForm />
    </main>
  );
};

export default Login;
