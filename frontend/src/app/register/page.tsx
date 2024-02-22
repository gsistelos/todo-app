import styles from "./page.module.css";
import Header from "../components/Header";
import RegisterForm from "../components/RegisterForm";

const Register = () => {
  return (
    <main className={styles.main}>
      <Header title="Register" subtitle="Create an account" />
      <RegisterForm />
    </main>
  );
};

export default Register;
