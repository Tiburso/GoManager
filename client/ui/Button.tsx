import { ReactNode } from 'react';

import style from '@/styles/Button.module.css';

export default function Button({
  children,
  onClick,
  className,
}: {
  children?: ReactNode;
  onClick?: (e: any) => void;
  className?: string;
}) {
  return (
    <button className={`${style.btn} ${className}`} onClick={onClick}>
      {children}
    </button>
  );
}
