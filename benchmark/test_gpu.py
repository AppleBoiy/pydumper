import torch

device = torch.device("cuda" if torch.cuda.is_available() else "cpu")
print(f"Using device: {device}")

# Example Model
model = torch.nn.Linear(10, 1).to(device)

# Example Tensor
data = torch.randn(5, 10).to(device)
output = model(data)

print(f"Output on {device}: {output}")