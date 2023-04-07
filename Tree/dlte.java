public class Delection {
	Node root;
	public class Node{
		Node left;
		Node right;
		int data;
		public Node(int data) {
			this.data=data;
		}
	}
	public void insert(int data) {
		root=insert(root, data);
	}
	public Node insert(Node root,int data) {
		Node newnode=new Node(data);
		if(root==null) {
			root=newnode;
			return root;
		}
		if(data<root.data) {
			root.left=insert(root.left, data);
		}else {
			root.right=insert(root.right, data);
		}
		return root;
	}
	public void delete(int key) {
		root=delete(root, key);
	}
	public Node delete(Node root,int key) {
		if(root==null) {
			return null;
		}
		if(key<root.data) {
			root.left=delete(root.left,key);
		}else if(key>root.data) {
			root.right=delete(root.right,key);
		}else {
			if(root.left==null) {
				return root.right;
			}else if(root.right==null) {
				return root.right;
			}
			root.data=findmin(root.right);
			root.right=delete(root.right,root.data);
		}
		return root;
	}
	public int findmin(Node node) {
		int min=node.data;
		while(node.left!=null) {
			min=node.left.data;
			node=node.left;
		}
		return min;
	}
	public void inorder(Node root) {
		if(root==null) {
			return;
		}
		inorder(root.left);
		System.out.print(root.data+" ");
		inorder(root.right);
	}
	public static void main(String[] args) {
		Delection list=new Delection();
		list.insert(48);
		list.insert(42);
		list.insert(19);
		list.insert(38);
		list.insert(49);
		list.insert(84);
		list.insert(52);
		list.insert(91);
		System.out.print("Before delection:-");
		list.inorder(list.root);
		System.out.println();
		System.out.print("After delection :-");
		list.delete(91);
		list.inorder(list.root);
	}

}