type Props = {
  error?: string;
  placeholder: string;
  type: string;
  name: string;
  onKeyDown?: (e: React.KeyboardEvent<HTMLInputElement>) => void;
  ref?: React.RefObject<HTMLInputElement>;
};

const FormInput = ({
  error,
  placeholder,
  type,
  name,
  onKeyDown,
  ref
}: Props) => {
  return (
    <div className="flex flex-col gap-1">
      {error && <span className="block text-red text-sm">{error}</span>}
      <input
        className="flex flex-col bg-secondary p-2 border border-contrast"
        placeholder={placeholder}
        type={type}
        name={name}
        onKeyDown={onKeyDown}
        ref={ref}
      />
    </div>
  );
};

export default FormInput;
