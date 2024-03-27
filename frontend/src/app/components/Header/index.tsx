'use client';

import Image from 'next/image';

import { useTheme } from '@/app/contexts/Theme';

import ThemeSwitcher from '../ThemeSwitcher';
import UserArea from '../UserArea';

type Props = {
  title: string;
};

const Header = ({ title }: Props) => {
  const { theme } = useTheme();

  const src = theme === 'dark' ? '/light-logo.png' : '/dark-logo.png';

  return (
    <header className="flex items-center justify-between p-6 border-b border-contrast">
      <div className="flex items-center gap-4">
        <Image width={24} height={24} src={src} alt="Logo" />
        <h1 className="text-2xl font-bold">{title}</h1>
      </div>
      <div className="flex items-center gap-14">
        <nav>
          <UserArea />
        </nav>
        <ThemeSwitcher />
      </div>
    </header>
  );
};

export default Header;
