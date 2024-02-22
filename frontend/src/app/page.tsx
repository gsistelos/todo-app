import styles from "./page.module.css";
import Header from "./components/Header";
import UserActions from "./components/UserActions";

const Home = () => {
  return (
    <main className={styles.main}>
      <Header title="Home" subtitle="todo-app" />
      <UserActions />
    </main>
  );
};

export default Home;
