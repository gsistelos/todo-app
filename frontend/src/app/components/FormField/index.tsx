import { useTheme } from "../../contexts/Theme";

type Props = {
  className?: string;
  label: string;
  type: string;
  name: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

const FormField = ({ className, label, type, name, onChange }: Props) => {
  const { textColor, bgSecondaryColor, borderColor } = useTheme();

  return (
    <div className={className}>
      <label className="font-medium">{label}</label>
      <input
        className={`p-1.5 ${textColor} ${bgSecondaryColor} border ${borderColor}`}
        type={type}
        name={name}
        onChange={onChange}
      />
    </div>
  );
};

export default FormField;
