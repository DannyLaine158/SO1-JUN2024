#include <linux/module.h>
// para usar KERN_INFO
#include <linux/kernel.h>
// para info_ram
#include <linux/sched.h>

//Header para los macros module_init y module_exit
#include <linux/init.h>
//Header necesario porque se usara proc_fs
#include <linux/proc_fs.h>
/* for copy_from_user */
#include <asm/uaccess.h>
/* Header para usar la lib seq_file y manejar el archivo en /proc*/
#include <linux/seq_file.h>
// para get_mm_rss
#include <linux/mm.h>

struct task_struct *cpu; // Estructura que almacena info del cpu

// Almacena los procesos
struct list_head *lstProcess;
// Estructura que almacena info de los procesos hijos
struct task_struct *child;
unsigned long rss;

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Modulo de CPU para el Lab de Sopes 1");
MODULE_AUTHOR("Dani :)");

static int escribir_archivo(struct seq_file *archivo, void *v) {
    for_each_process(cpu) {
        seq_printf(archivo, "PID%d", cpu->pid);
        seq_printf(archivo, ",");
        seq_printf(archivo, "%s", cpu->comm);
        seq_printf(archivo, ",");
        seq_printf(archivo, "%lu", cpu->__state);
        seq_printf(archivo, ",");

        if (cpu->mm) {
            rss = get_mm_rss(cpu->mm) << PAGE_SHIFT;
            seq_printf(archivo, "%lu", rss);
        } else {
            seq_printf(archivo, "%s", "");
        }
        seq_printf(archivo, ",");

        seq_printf(archivo, "%d", cpu->cred->user->uid);
        seq_printf(archivo, ",");

        list_for_each(lstProcess, &(cpu->children)) {
            child = list_entry(lstProcess, struct task_struct, sibling);
            seq_printf(archivo, "Child:%d", child->pid);
            seq_printf(archivo, ".");
            seq_printf(archivo, "%s", child->comm);
            seq_printf(archivo, ".");
            seq_printf(archivo, "%d", child->__state);
            seq_printf(archivo, ".");

             if (child->mm) {
                rss = get_mm_rss(child->mm) << PAGE_SHIFT;
                seq_printf(archivo, "%lu", rss);
            } else {
                seq_printf(archivo, "%s", "");
            }
            seq_printf(archivo, ".");

            seq_printf(archivo, "%d", child->cred->user->uid);
        }
    }

    return 0;
}

//Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
static int al_abrir(struct inode *inode, struct file *file)
{
    return single_open(file, escribir_archivo, NULL);
}

//Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operaciones =
{
    .proc_open = al_abrir,
    .proc_read = seq_read
};

//Funcion a ejecuta al insertar el modulo en el kernel con insmod
static int _insert(void)
{
    proc_create("cpu_201800722", 0, NULL, &operaciones);
    printk(KERN_INFO "Jose Daniel Velasquez Orozco\n");
    return 0;
}

//Funcion a ejecuta al remover el modulo del kernel con rmmod
static void _remove(void)
{
    remove_proc_entry("cpu_201800722", NULL);
    printk(KERN_INFO "Segundo Semestre 2023\n");
}

module_init(_insert);
module_exit(_remove);
