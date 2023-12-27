import Button from '@/ui/Button';

import { useState } from 'react';

import { Company } from '@/lib/types';

import style from '@/styles/company/CompanyForm.module.css';

export default function CompanyForm({
  onClose,
  onSave,
}: {
  onClose: () => void;
  onSave: (e: any) => void;
}) {
  const [name, setName] = useState('');
  const [candidatePortal, setCandidatePortal] = useState('');

  const handleCloseClick = (e: MouseEvent) => {
    e.preventDefault();
    onClose();
  };

  const handleSaveClick = (e: MouseEvent) => {
    // validate the values of the name and candidate_portal
    e.preventDefault();
    const c: Company = {
      name,
      company_portal: candidatePortal,
    };

    // pass the company values to the save function
    onSave(c);
  };

  return (
    <div>
      <div className={style.form}>
        <label htmlFor="name">Name</label>
        <input
          type="text"
          id="name"
          name="name"
          placeholder="Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />
        <label htmlFor="candidate_portal">Candidate Portal</label>
        <input
          type="url"
          id="candidate_portal"
          name="candidate_portal"
          placeholder="Candidate portal"
          value={candidatePortal}
          onChange={(e) => setCandidatePortal(e.target.value)}
        />
      </div>
      {/* buttons */}
      <div className={style.btns}>
        <Button className={style.cancel} onClick={handleCloseClick}>
          Close
        </Button>
        <Button onClick={handleSaveClick}>Add</Button>
      </div>
    </div>
  );
}
