import { ReactNode } from 'react';

import style from '@/app/styles/Button.module.css';

export default function Button({
  children,
  onClick,
}: {
  children?: ReactNode;
  onClick?: () => void;
}) {
  return (
    <button className={style.btn} onClick={onClick}>
      {children}
    </button>
  );
}
