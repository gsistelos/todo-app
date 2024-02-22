import styles from "./FormGroup.module.css";

type Props = {
  label: string;
  type: string;
  name: string;
  value: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
  error: string;
};

const FormGroup = ({ label, type, name, value, onChange, error }: Props) => {
  return (
    <div className={styles.formGroup}>
      <label>{label}</label>
      <input type={type} name={name} value={value} onChange={onChange} />
      {error && <p className={styles.error}>{error}</p>}
    </div>
  );
};

export default FormGroup;
