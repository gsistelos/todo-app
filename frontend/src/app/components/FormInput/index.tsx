'use client';

import { useTheme } from '@/app/contexts/Theme';

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
  const { textColor, bgSecondaryColor, borderColor } = useTheme();

  return (
    <div className="flex flex-col gap-1">
      {error && <span className="block text-red-500 text-sm">{error}</span>}
      <input
        className={`flex flex-col ${bgSecondaryColor} p-2 border ${borderColor} ${textColor}`}
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
