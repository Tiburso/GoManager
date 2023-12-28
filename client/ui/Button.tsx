import { ReactNode } from 'react';

import style from '@/styles/Button.module.css';

type ButtonType = 'button' | 'submit' | 'reset';

export default function Button({
  type = 'button',
  disabled,
  children,
  onClick,
  className,
}: {
  type?: ButtonType;
  children?: ReactNode;
  disabled?: boolean;
  onClick?: (e: any) => void;
  className?: string;
}) {
  return (
    <button
      type={type}
      className={`${style.btn} ${className}`}
      onClick={onClick}
      aria-disabled={disabled}
    >
      {children}
    </button>
  );
}
