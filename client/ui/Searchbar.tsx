import style from '@/styles/Searchbar.module.css';
import { FontAwesomeIcon } from '@fortawesome/react-fontawesome';
import { faSearch } from '@fortawesome/free-solid-svg-icons';

export default function Searchbar({
  type,
  placeholder,
  value,
  onChange,
}: {
  type: string;
  placeholder: string;
  value?: string;
  onChange?: (e: any) => void;
}) {
  return (
    <div className={style.searchbar}>
      <FontAwesomeIcon icon={faSearch} color="lightgray" />
      <input
        type={type}
        placeholder={placeholder}
        value={value}
        onChange={onChange}
      />
    </div>
  );
}
