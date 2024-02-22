import styles from "./Header.module.css";

type Props = {
  title: string;
  subtitle?: string;
};

const Header = ({ title, subtitle }: Props) => {
  return (
    <header className={styles.header}>
      <h1 className={styles.title}>{title}</h1>
      {subtitle && <h2 className={styles.subtitle}>{subtitle}</h2>}
    </header>
  );
};

export default Header;
