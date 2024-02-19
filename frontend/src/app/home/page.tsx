import styles from "./page.module.css";

const Home: React.FC = () => {
  return (
    <main className={styles.main}>
      <h1 className={styles.title}>Home</h1>
      <p className={styles.paragraph}>This is the Home page</p>
    </main>
  );
};

export default Home;
