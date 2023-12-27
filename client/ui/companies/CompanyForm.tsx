import Button from '@/ui/Button';

import style from '@/styles/company/CompanyForm.module.css';

export default function CompanyForm({
  onClose,
  handleSaveClick,
}: {
  onClose: () => void;
  handleSaveClick: () => void;
}) {
  const handleCloseClick = (e: MouseEvent) => {
    e.preventDefault();
    onClose();
  };

  // Here after I handle the save click I should send the data to the server
  // And also update the UI

  return (
    <div>
      <div className={style.form}>
        <label htmlFor="name">Name</label>
        <input type="text" id="name" name="name" />
        <label htmlFor="candidate_portal">Candidate Portal</label>
        <input type="text" id="candidate_portal" name="candidate_portal" />
      </div>
      {/* buttons */}
      <div className={style.btns}>
        <Button className={style.cancel} onClick={handleCloseClick}>
          Close
        </Button>
        <Button onClick={handleSaveClick}>Save</Button>
      </div>
    </div>
  );
}
