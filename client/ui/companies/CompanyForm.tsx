import { Company } from '@/lib/types';
import { addCompany } from '@/lib/companies';
import Button from '@/ui/Button';

import { useFormStatus, useFormState } from 'react-dom';

import style from '@/styles/company/CompanyForm.module.css';

export default function CompanyForm({
  onClose,
  company,
}: {
  onClose: () => void;
  company?: Company;
}) {
  const { pending } = useFormStatus();

  const [state, formAction] = useFormState(
    async (previousState: any, formData: FormData) => {
      try {
        await addCompany(formData);
        onClose();
        return {
          message: 'Company added',
        };
      } catch (error) {
        console.error(error);
        return {
          message: 'Error adding company',
        };
      }
    },
    {
      message: null as null | string,
    },
  );

  return (
    <div>
      <form action={formAction}>
        <div className={style.form}>
          <label htmlFor="name">Name</label>
          <input
            required
            type="text"
            id="name"
            name="name"
            placeholder="Name"
            value={company?.name}
          />
          <label htmlFor="candidate_portal">Candidate Portal</label>
          <input
            required
            type="url"
            id="candidate_portal"
            name="candidate_portal"
            placeholder="Candidate portal"
            value={company?.candidate_portal}
          />
        </div>
        <div className={style.btns}>
          <Button className={style.cancel} onClick={onClose}>
            Close
          </Button>
          <Button type="submit" disabled={pending}>
            Add
          </Button>
          {/* it should be the state.message here that would be returned from the server actions this is only for screen readers*/}
          <p aria-live="polite" className="sr-only" role="status">
            {state?.message}
          </p>
        </div>
      </form>
    </div>
  );
}
