import { useTheme } from "../../contexts/Theme";

type Props = {
  className?: string;
  label: string;
  type: string;
  name: string;
  onChange: (e: React.ChangeEvent<HTMLInputElement>) => void;
};

const FormField = ({ className, label, type, name, onChange }: Props) => {
  const { lights } = useTheme();

  const input = lights
    ? "text-black border-black bg-gray-200"
    : "text-white border-white bg-gray-800";

  return (
    <div className={className}>
      <label className="font-medium">{label}</label>
      <input
        className={`p-1.5 border ${input}`}
        type={type}
        name={name}
        onChange={onChange}
      />
    </div>
  );
};

export default FormField;
