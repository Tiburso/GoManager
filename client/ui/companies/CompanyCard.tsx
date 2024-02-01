import { useState } from 'react';

import { Company } from '@/lib/types';

import Button from '@/ui/Button';
import Modal from '@/ui/Modal';
import CompanyForm from '@/ui/companies/CompanyForm';
import Link from 'next/link';

import style from '@/styles//company/CompanyCard.module.css';

export default function CompanyCard({ company }: { company: Company }) {
  const [showModal, setShowModal] = useState(false);

  return (
    <div className={style.card}>
      <Link
        className={style.link}
        href={`/companies/${company.id}`}
        key={company.id}
      >
        <h2>{company.name}</h2>
        <p>{company.candidate_portal}</p>
      </Link>

      {/* here i will have two buttons that will show up as i hover over the card */}
      {
        <div className={style.btns}>
          <Button className={style.delete} onClick={() => console.log('delete')}>Delete</Button>

          {/* this button should be a form hidden inputs */}
          <Button onClick={() => setShowModal(true)} type="submit">
            Edit
          </Button>
        </div>
      }
      {showModal && (
        <Modal onClose={() => setShowModal(false)} title="Update company">
          <CompanyForm onClose={() => setShowModal(false)} company={company} />
        </Modal>
      )}
    </div>
  );
}
