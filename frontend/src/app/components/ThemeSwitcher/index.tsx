'use client';

import Image from 'next/image';

import { useTheme } from '@/app/contexts/Theme';

const ThemeSwitcher = () => {
  const { theme, toggleTheme } = useTheme();

  const { src, alt } =
    theme === 'dark'
      ? {
        src: '/light.png',
        alt: 'Light icon',
      }
      : {
        src: '/dark.png',
        alt: 'Dark icon',
      };

  return (
    <button
      className="p-1 rounded-full hover:bg-secondary"
      onClick={() => toggleTheme()}
    >
      <Image width={24} height={24} src={src} alt={alt} />
    </button>
  );
};

export default ThemeSwitcher;
